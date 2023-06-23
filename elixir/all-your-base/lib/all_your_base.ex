defmodule AllYourBase do
  @doc """
  Given a number in input base, represented as a sequence of digits, converts it to output base,
  or returns an error tuple if either of the bases are less than 2
  """

  @spec convert(list, integer, integer) :: {:ok, list} | {:error, String.t()}
  def convert(digits, input_base, output_base) do
    cond do
      output_base < 2 -> {:error, "output base must be >= 2"}
      input_base < 2 -> {:error, "input base must be >= 2"}
      Enum.any?(digits, &(&1 < 0 || &1 >= input_base)) -> {:error, "all digits must be >= 0 and < input base"}
      true -> {:ok, find_digits(base_ten(digits, input_base, 0), output_base, [])}
    end
  end

  defp base_ten([], _, prev), do: prev
  defp base_ten(_digits = [head | tail], input_base, prev) do
    base_ten(tail, input_base, input_base * prev + head)
  end

  defp find_digits(0, _, []), do: [0]
  defp find_digits(0, _, output_digits), do: output_digits
  defp find_digits(number, output_base, output_digits) do
    find_digits(div(number, output_base), output_base, [rem(number, output_base) | output_digits]) 
  end
end
