defmodule Sublist do
  @doc """
  Returns whether the first list is a sublist or a superlist of the second list
  and if not whether it is equal or unequal to the second list.
  """
  def compare(a, b) do
    cond do
      a == b -> :equal
      length(a) < length(b) and sublist?(b, a) -> :sublist
      length(a) > length(b) and sublist?(a, b) -> :superlist
      true -> :unequal
    end
  end

  def sublist?([], _), do: false

  def sublist?(a = [_|t], b) do
    List.starts_with?(a,b) or sublist?(t, b)
  end
end
