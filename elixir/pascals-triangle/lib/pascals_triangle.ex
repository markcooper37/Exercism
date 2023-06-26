defmodule PascalsTriangle do
  @doc """
  Calculates the rows of a pascal triangle
  with the given height
  """
  @spec rows(integer) :: [[integer]]
  def rows(num) do
    construct_rows([[1]], num)
  end

  defp construct_rows(rows, req_rows) when length(rows) == req_rows, do: Enum.reverse(rows)
  defp construct_rows(prev_rows = [last_row | _], req_rows) do
    construct_rows([create_row(last_row) | prev_rows], req_rows)
  end

  defp create_row(prev_row) do
    Enum.map(Enum.zip([[0 | prev_row], prev_row]), fn {x1, x2} -> x1 + x2 end) ++ [1]
  end
end
