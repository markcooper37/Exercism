defmodule Sieve do
  @doc """
  Generates a list of primes up to a given limit.
  """
  @spec primes_to(non_neg_integer) :: [non_neg_integer]
  def primes_to(limit), do: sieve(Enum.to_list(2..limit//1), limit, [])

  defp sieve([], _, acc), do: acc
  defp sieve([prime | numbers], limit, primes) do
    sieve(numbers -- Enum.to_list(prime..limit//prime), limit, primes ++ [prime])
  end
end
