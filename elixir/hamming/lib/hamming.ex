defmodule Hamming do
  @doc """
  Returns number of differences between two strands of DNA, known as the Hamming Distance.

  ## Examples

  iex> Hamming.hamming_distance('AAGTCATA', 'TAGCGATC')
  {:ok, 4}
  """
  @spec hamming_distance([char], [char]) :: {:ok, non_neg_integer} | {:error, String.t()}
  def hamming_distance(strand1, strand2) do
    cond do
      length(strand1) != length(strand2) -> {:error, "strands must be of equal length"}
      true ->{:ok, Enum.reduce(0..length(strand1)-1, 0, fn(index, acc) -> if Enum.at(strand1, index) == Enum.at(strand2, index), do: acc, else: acc + 1 end)}
    end
  end
end
