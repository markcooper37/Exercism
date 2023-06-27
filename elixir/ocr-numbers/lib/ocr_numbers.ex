defmodule OcrNumbers do
  @numbers %{{' _ ', '| |', '|_|', '   '} => "0", {'   ', '  |', '  |', '   '} => "1", {' _ ', ' _|', '|_ ', '   '} => "2", {' _ ', ' _|', ' _|', '   '} => "3", {'   ', '|_|', '  |', '   '} => "4", {' _ ', '|_ ', ' _|', '   '} => "5", {' _ ', '|_ ', '|_|', '   '} => "6", {' _ ', '  |', '  |', '   '} => "7", {' _ ', '|_|', '|_|', '   '} => "8", {' _ ', '|_|', ' _|', '   '} => "9"}

  @doc """
  Given a 3 x 4 grid of pipes, underscores, and spaces, determine which number is represented, or
  whether it is garbled.
  """
  @spec convert([String.t()]) :: {:ok, String.t()} | {:error, String.t()}
  def convert(input) do
    cond do
      rem(length(input), 4) != 0 -> {:error, "invalid line count"}
      Enum.any?(input, &(rem(String.length(&1), 3) != 0)) -> {:error, "invalid column count"}
      true -> {:ok, do_convert(input)}
    end
  end

  defp do_convert(input) do
    input
    |> Enum.chunk_every(4)
    |> Enum.map(&construct_line/1)
    |> Enum.join(",")
  end

  defp construct_line(input) do
    input
    |> Enum.map(&String.to_charlist/1)
    |> Enum.map(&Enum.chunk_every(&1, 3))
    |> Enum.zip()
    |> Enum.map(&(if (num = @numbers[&1]) != nil, do: num, else: "?"))
    |> Enum.join("")
  end
end
