defmodule Garden do
  @default_names [:alice, :bob, :charlie, :david, :eve, :fred, :ginny, :harriet, :ileana, :joseph, :kincaid, :larry]
  @plants %{"V" => :violets, "R" => :radishes, "G" => :grass, "C" => :clover}

  @doc """
    Accepts a string representing the arrangement of cups on a windowsill and a
    list with names of students in the class. The student names list does not
    have to be in alphabetical order.

    It decodes that string into the various gardens for each student and returns
    that information in a map.
  """

  @spec info(String.t(), list) :: map
  def info(info_string, student_names \\ @default_names) do
    class_map = Enum.reduce(student_names, %{}, fn name, acc -> Map.put(acc, name, {}) end)
    info_string
    |> String.split("\n", trim: true)
    |> Enum.map(&String.split(&1, "", trim: true))
    |> Enum.map(&Enum.chunk_every(&1, 2))
    |> Enum.zip
    |> Enum.map(fn {row_1, row_2} -> row_1 ++ row_2 end)
    |> Enum.zip(Enum.sort(student_names))
    |> Enum.reduce(class_map, fn {letters, name}, acc -> Map.put(acc, name, get_plants(letters)) end)
  end

  defp get_plants(letters) do
    letters
    |> Enum.map(&@plants[&1])
    |> List.to_tuple
  end
end
