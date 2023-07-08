defmodule Queens do
  @type t :: %Queens{black: {integer, integer}, white: {integer, integer}}
  defstruct [:white, :black]

  @doc """
  Creates a new set of Queens
  """
  @spec new(Keyword.t()) :: Queens.t()
  def new(opts \\ []) do
    cond do
      !Enum.all?(opts, fn {colour, position} -> valid_opt?(colour, position) end) -> raise ArgumentError
      opts[:white] != nil and opts[:white] == opts[:black] -> raise ArgumentError
      true -> %Queens{white: opts[:white], black: opts[:black]}
    end
  end

  defp valid_opt?(colour, {x, y}) do
    colour in [:white, :black] && x in 0..7 && y in 0..7
  end

  @doc """
  Gives a string representation of the board with
  white and black queen locations shown
  """
  @spec to_string(Queens.t()) :: String.t()
  def to_string(queens) do
    0..7
    |> Enum.map(&make_row(&1, queens))
    |> Enum.map(&Enum.join(&1, " "))
    |> Enum.join("\n")
  end

  defp make_row(row, queens), do: 0..7 |> Enum.map(&value_at_pos(row, &1, queens))

  defp value_at_pos(row, column, queens) do
    cond do
      queens.black == {row, column} -> "B"
      queens.white == {row, column} -> "W"
      true -> "_"
    end
  end

  @doc """
  Checks if the queens can attack each other
  """
  @spec can_attack?(Queens.t()) :: boolean
  def can_attack?(queens) when queens.white == nil or queens.black == nil, do: false
  def can_attack?(%Queens{white: {x1, y1}, black: {x2, y2}}) do
    x1 == x2 || y1 == y2 || x1 + y1 == x2 + y2 || x1 - y1 == x2 - y2
  end
end
