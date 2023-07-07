defmodule Luhn do
  @doc """
  Checks if the given number is valid via the luhn formula
  """
  @spec valid?(String.t()) :: boolean
  def valid?(number) do
    number = String.replace(number, " ", "")
    cond do
      Regex.match?(~r/\p{P}|[a-zA-z]/, number) -> false
      number == "0" -> false
      true -> rem(luhn_value(number), 10) == 0
    end
  end

  defp luhn_value(number) do
    number
    |> String.split("", trim: true)
    |> Enum.map(&String.to_integer/1)
    |> Enum.reverse
    |> Enum.with_index
    |> Enum.map(fn {digit, index} -> if rem(index, 2) == 1, do: new_value(digit), else: digit end)
    |> Enum.sum
  end

  defp new_value(digit) do
    if digit >= 5, do: digit * 2 - 9, else: digit * 2
  end
end
