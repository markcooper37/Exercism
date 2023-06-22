defmodule FoodChain do
  @creatures %{1 => "fly", 2 => "spider", 3 => "bird", 4 => "cat", 5 => "dog", 6 => "goat", 7 => "cow", 8 => "horse"}

  @lines %{
    1 => "",
    2 => "It wriggled and jiggled and tickled inside her.\n",
    3 => "How absurd to swallow a bird!\n",
    4 => "Imagine that, to swallow a cat!\n",
    5 => "What a hog, to swallow a dog!\n",
    6 => "Just opened her throat and swallowed a goat!\n",
    7 => "I don't know how she swallowed a cow!\n",
    8 => ""
  }

  @spider_line " that wriggled and jiggled and tickled inside her"

  @doc """
  Generate consecutive verses of the song 'I Know an Old Lady Who Swallowed a Fly'.
  """
  @spec recite(start :: integer, stop :: integer) :: String.t()
  def recite(start, stop) when start > stop, do: ""
  def recite(start, stop) do
    "#{verse(start)}#{if start != stop, do: "\n"}#{recite(start + 1, stop)}"
  end

  defp verse(number) do
    "#{beginning(number)}#{middle(number)}#{ending(number)}"
  end

  defp beginning(verse) do
    "I know an old lady who swallowed a #{@creatures[verse]}.\n#{@lines[verse]}"
  end

  defp middle(verse) when verse in [1, 8], do: ""
  defp middle(verse) do
    "She swallowed the #{@creatures[verse]} to catch the #{@creatures[verse-1]}#{if verse == 3, do: @spider_line}.\n#{middle(verse-1)}"
  end

  defp ending(8), do: "She's dead, of course!\n"
  defp ending(_), do: "I don't know why she swallowed the fly. Perhaps she'll die.\n"
end
