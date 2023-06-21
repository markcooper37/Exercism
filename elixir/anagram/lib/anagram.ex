defmodule Anagram do
  @doc """
  Returns all candidates that are anagrams of, but not equal to, 'base'.
  """
  @spec match(String.t(), [String.t()]) :: [String.t()]
  def match(base, candidates) do
    base_letters = count_letters(String.downcase(base))
    Enum.filter(candidates, fn (candidate) -> String.downcase(base) != String.downcase(candidate) and base_letters == count_letters(String.downcase(candidate)) end)
  end

  defp count_letters(word) do
    Enum.reduce(String.graphemes(word), %{}, fn char, acc -> Map.put(acc, char, (acc[char] || 0) + 1) end)
  end
end
