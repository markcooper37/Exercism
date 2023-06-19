defmodule DNA do
  def encode_nucleotide(code_point) do
    case code_point do
      ?\s -> 0b0000
      ?A -> 0b0001
      ?C -> 0b0010
      ?G -> 0b0100
      ?T -> 0b1000
    end
  end

  def decode_nucleotide(encoded_code) do
    case encoded_code do
      0b0000 -> ?\s
      0b0001 -> ?A
      0b0010 -> ?C
      0b0100 -> ?G
      0b1000 -> ?T
    end
  end

  def encode(dna) do
    do_encode(dna, <<>>)
  end

  defp do_encode('', encoding), do: encoding
  defp do_encode([head | tail], encoding), do: do_encode(tail, <<encoding::bitstring, encode_nucleotide(head)::4>>)

  def decode(dna) do
    do_decode(dna, '')
  end

  defp do_decode(<<0::0>>, decoding), do: decoding
  defp do_decode(<<value::4, rest::bitstring>>, decoding), do: do_decode(rest, decoding ++ [decode_nucleotide(value)])
end
