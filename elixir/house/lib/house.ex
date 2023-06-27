defmodule House do
  @actions %{2 => "lay in", 3 => "ate", 4 => "killed", 5 => "worried", 6 => "tossed", 7 => "milked", 8 => "kissed", 9 => "married", 10 => "woke", 11 => "kept", 12 => "belonged to"}
  @objects %{1 => "house that Jack built", 2 => "malt", 3 => "rat", 4 => "cat", 5 => "dog", 6 => "cow with the crumpled horn", 7 => "maiden all forlorn", 8 => "man all tattered and torn", 9 => "priest all shaven and shorn", 10 => "rooster that crowed in the morn", 11 => "farmer sowing his corn", 12 => "horse and the hound and the horn"}

  @doc """
  Return verses of the nursery rhyme 'This is the House that Jack Built'.
  """
  @spec recite(start :: integer, stop :: integer) :: String.t()
  def recite(start, stop) do
    start..stop
    |> Enum.map(&construct_verse/1)
    |> Enum.join("")
  end

  defp construct_verse(verse), do: "This is #{construct_middle(verse)}.\n"

  defp construct_middle(1), do: "the #{@objects[1]}"
  defp construct_middle(verse) do
    "the #{@objects[verse]} that #{@actions[verse]} " <> construct_middle(verse - 1)
  end
end
