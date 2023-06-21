defmodule Proverb do
  @doc """
  Generate a proverb from a list of strings.
  """
  @spec recite(strings :: [String.t()]) :: String.t()
  def recite([]), do: ""
  def recite(strings = [head | tail]) do
    Enum.reduce(Enum.zip(strings, tail), "", fn ({first, second}, acc) -> acc <> "For want of a #{first} the #{second} was lost.\n" end) <> "And all for the want of a #{head}.\n"
  end
end
