defmodule Series do
  @doc """
  Finds the largest product of a given number of consecutive numbers in a given string of numbers.
  """
  @spec largest_product(String.t(), non_neg_integer) :: non_neg_integer
  def largest_product(number_string, size) do
    cond do
      String.length(number_string) < size || size <= 0 -> raise ArgumentError
      true -> find_largest_product(number_string, size)
    end
  end

  defp find_largest_product(number_string, size) do
    number_string
    |> String.split("", trim: true)
    |> Enum.map(&String.to_integer/1)
    |> Enum.chunk_every(size, 1, :discard)
    |> Enum.map(&Enum.product/1)
    |> Enum.max
  end
end
