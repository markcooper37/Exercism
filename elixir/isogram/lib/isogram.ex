defmodule Isogram do
  @doc """
  Determines if a word or sentence is an isogram
  """
  @spec isogram?(String.t()) :: boolean
  def isogram?(sentence), do: !Regex.match?(~r/([a-z]).*\1/, String.downcase(sentence))
end
