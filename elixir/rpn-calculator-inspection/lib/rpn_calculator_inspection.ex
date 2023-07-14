defmodule RPNCalculatorInspection do
  def start_reliability_check(calculator, input) do
    %{input: input, pid: spawn_link(fn -> calculator.(input) end)}
  end

  def await_reliability_check_result(%{pid: pid, input: input}, results) do
    receive do
      {:EXIT, ^pid, :normal} -> Map.put(results, input, :ok)
      {:EXIT, ^pid, _} -> Map.put(results, input, :error)
    after
      100 -> Map.put(results, input, :timeout)
    end
  end

  def reliability_check(calculator, inputs) do
    old_flag_value = Process.flag(:trap_exit, true)
    output = inputs
    |> Enum.map(fn input -> start_reliability_check(calculator, input) end)
    |> Enum.map(fn element -> await_reliability_check_result(element, %{}) end)
    |> Enum.reduce(%{}, fn result, acc -> Map.merge(acc, result) end)
    Process.flag(:trap_exit, old_flag_value)
    output
  end

  def correctness_check(calculator, inputs) do
    inputs
    |> Enum.map(fn input -> Task.async(fn -> calculator.(input) end) end)
    |> Enum.map(fn task -> Task.await(task, 100) end)
  end
end
