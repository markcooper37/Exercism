defmodule ResistorColorTrio do
  @values %{black: 0, brown: 1, red: 2, orange: 3, yellow: 4, green: 5, blue: 6, violet: 7, grey: 8, white: 9}

  @doc """
  Calculate the resistance value in ohms from resistor colors
  """
  @spec label(colors :: [atom]) :: {number, :ohms | :kiloohms | :megaohms | :gigaohms}
  def label(colors) do
    value = (@values[Enum.at(colors, 0)] * 10 + @values[Enum.at(colors, 1)]) * (10 ** @values[Enum.at(colors, 2)])
    cond do
      value >= 1000000000 -> {value / 1000000000, :gigaohms}
      value >= 1000000 -> {value / 1000000, :megaohms}
      value >= 1000 -> {value / 1000, :kiloohms}
      true -> {value, :ohms}
    end
  end
end
