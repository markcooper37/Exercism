defmodule SumOfMultiples do
  @doc """
  Adds up all numbers from 1 to a given end number that are multiples of the factors provided.
  """
  @spec to(non_neg_integer, [non_neg_integer]) :: non_neg_integer
  def to(limit, factors) do
    factors
    |> Enum.map(&multiples(limit, &1))
    |> Enum.concat
    |> Enum.uniq
    |> Enum.sum
  end

  defp multiples(_, 0), do: []
  defp multiples(limit, factor), do: factor..(limit-1)//factor
end
