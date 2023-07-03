defmodule Acronym do
  @doc """
  Generate an acronym from a string.
  "This is a string" => "TIAS"
  """
  @spec abbreviate(String.t()) :: String.t()
  def abbreviate(string) do
    string
    |> String.replace("-", " ")
    |> String.replace(~r/\p{P}/, "")
    |> String.split(~r/\s+/)
    |> Enum.map_join("", fn (word) -> String.upcase(String.at(word, 0)) end)
  end
end
