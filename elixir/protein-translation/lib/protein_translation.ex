defmodule ProteinTranslation do
  @doc """
  Given an RNA string, return a list of proteins specified by codons, in order.
  """
  @spec of_rna(String.t()) :: {:ok, list(String.t())} | {:error, String.t()}
  def of_rna(rna) do
    rna
    |> String.codepoints()
    |> Enum.chunk_every(3)
    |> Enum.map(&Enum.join/1)
    |> construct_proteins
  end

  defp construct_proteins(codons, proteins \\ [])
  defp construct_proteins([], proteins), do: {:ok, proteins}
  defp construct_proteins([codon | codons], proteins) do
    case of_codon(codon) do
      {:error, _} -> {:error, "invalid RNA"}
      {:ok, "STOP"} -> {:ok, proteins}
      {:ok, protein} -> construct_proteins(codons, proteins ++ [protein])
    end
  end

  @doc """
  Given a codon, return the corresponding protein

  UGU -> Cysteine
  UGC -> Cysteine
  UUA -> Leucine
  UUG -> Leucine
  AUG -> Methionine
  UUU -> Phenylalanine
  UUC -> Phenylalanine
  UCU -> Serine
  UCC -> Serine
  UCA -> Serine
  UCG -> Serine
  UGG -> Tryptophan
  UAU -> Tyrosine
  UAC -> Tyrosine
  UAA -> STOP
  UAG -> STOP
  UGA -> STOP
  """
  @spec of_codon(String.t()) :: {:ok, String.t()} | {:error, String.t()}
  def of_codon(codon) do
    case codon do
      "AUG" -> {:ok, "Methionine"}
      codon when codon in ["UUU", "UUC"] -> {:ok, "Phenylalanine"}
      codon when codon in ["UUA", "UUG"] -> {:ok, "Leucine"}
      codon when codon in ["UCU", "UCC", "UCA", "UCG"] -> {:ok, "Serine"}
      codon when codon in ["UAU", "UAC"] -> {:ok, "Tyrosine"}
      codon when codon in ["UGU", "UGC"] -> {:ok, "Cysteine"}
      "UGG" -> {:ok, "Tryptophan"}
      codon when codon in ["UAA", "UAG", "UGA"] -> {:ok, "STOP"}
      _ -> {:error, "invalid codon"}
    end
  end
end
