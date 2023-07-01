defmodule WordCount do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t()) :: map
  def count(sentence) do
    Regex.scan(~r/\b\w+['-]?\w*\b/u, String.replace(String.downcase(sentence), "_", " "))
    |> List.flatten
    |> Enum.frequencies
  end
end
