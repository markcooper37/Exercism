defmodule PalindromeProducts do
  @doc """
  Generates all palindrome products from an optionally given min factor (or 1) to a given max factor.
  """
  @spec generate(non_neg_integer, non_neg_integer) :: map
  def generate(max_factor, min_factor \\ 1) 
  def generate(max_factor, min_factor) when min_factor > max_factor, do: raise ArgumentError
  def generate(max_factor, min_factor) do
    for x <- min_factor..max_factor, y <- x..max_factor, palindrome?(x*y) do
      {x*y, [x, y]}
    end
    |> Enum.reduce(%{}, &combine_factors/2)
    |> Map.new
  end

  defp combine_factors({number, factors}, factor_map) do
    Map.update(factor_map, number, [factors], &(&1 ++ [factors]))
  end

  defp palindrome?(number) do
    Integer.to_string(number) == String.reverse(Integer.to_string(number))
  end
end
