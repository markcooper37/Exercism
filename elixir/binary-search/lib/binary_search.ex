defmodule BinarySearch do
  @doc """
    Searches for a key in the tuple using the binary search algorithm.
    It returns :not_found if the key is not in the tuple.
    Otherwise returns {:ok, index}.

    ## Examples

      iex> BinarySearch.search({}, 2)
      :not_found

      iex> BinarySearch.search({1, 3, 5}, 2)
      :not_found

      iex> BinarySearch.search({1, 3, 5}, 5)
      {:ok, 2}

  """

  @spec search(tuple, integer) :: {:ok, integer} | :not_found
  def search({}, _key) do
    :not_found
  end

  def search(numbers, key) do
    perform_search(numbers, key, 0, tuple_size(numbers) - 1) 
  end

  defp perform_search(_numbers, _key, lower, upper) when lower > upper do
    :not_found
  end

  defp perform_search(numbers, key, lower, upper) do
    middle = div(upper + lower, 2)
    element = elem(numbers, middle)
    cond do
      element == key -> {:ok, middle}
      element < key -> perform_search(numbers, key, middle + 1, upper)
      element > key -> perform_search(numbers, key, lower, middle - 1)
    end
  end
end
