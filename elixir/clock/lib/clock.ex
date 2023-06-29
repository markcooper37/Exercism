defmodule Clock do
  defstruct hour: 0, minute: 0

  @doc """
  Returns a clock that can be represented as a string:

      iex> Clock.new(8, 9) |> to_string
      "08:09"
  """
  @spec new(integer, integer) :: Clock
  def new(hour, minute) do
    total_minutes = rem(rem(60 * hour + minute, 1440) + 1440, 1440)
    %Clock{hour: rem(div(total_minutes, 60), 24), minute: rem(total_minutes, 60)}
  end

  @doc """
  Adds two clock times:

      iex> Clock.new(10, 0) |> Clock.add(3) |> to_string
      "10:03"
  """
  @spec add(Clock, integer) :: Clock
  def add(%Clock{hour: hour, minute: minute}, add_minute), do: new(hour, minute + add_minute)

  defimpl String.Chars, for: Clock do
    def to_string(clock), do: "#{pad_time(clock.hour)}:#{pad_time(clock.minute)}"

    defp pad_time(time) do
      time
      |> Integer.to_string
      |> String.pad_leading(2, "0")
    end
  end
end
