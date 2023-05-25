defmodule Username do
  def sanitize([]), do: ''

  def sanitize([head | tail]) do
    case head do
      ?ä -> 'ae' ++ sanitize(tail)
      ?ö -> 'oe' ++ sanitize(tail)
      ?ü -> 'ue' ++ sanitize(tail)
      ?ß -> 'ss' ++ sanitize(tail)
      head when head >= ?a and head <= ?z or head == ?_ -> [head | sanitize(tail)]
      _ -> sanitize(tail)
    end
  end
end
