# Use the Plot struct as it is provided
defmodule Plot do
  @enforce_keys [:plot_id, :registered_to]
  defstruct [:plot_id, :registered_to]
end

defmodule CommunityGarden do
  def start(opts \\ []) do
    Agent.start(fn -> {%{}, 1} end, opts)
  end

  def list_registrations(pid) do
    Agent.get(pid, fn {plots, _} -> Map.values(plots) end)
  end

  def register(pid, register_to) do
    state = Agent.get(pid, fn {_, state} -> state end)
    plot = %Plot{plot_id: state, registered_to: register_to}
    Agent.update(pid, fn {plots, state} -> {Map.put(plots, state, plot), state + 1} end)
    plot
  end

  def release(pid, plot_id) do
    Agent.update(pid, fn {plots, state} -> {Map.delete(plots, plot_id), state} end)
  end

  def get_registration(pid, plot_id) do
    Agent.get(pid, fn {plots, _} -> Map.get(plots, plot_id) || {:not_found, "plot is unregistered"} end)
  end
end
