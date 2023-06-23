defmodule RomanNumerals do
  @units %{0 => "", 1 => "I", 2 => "II", 3 => "III", 4 => "IV", 5 => "V", 6 => "VI", 7 => "VII", 8 => "VIII", 9 => "IX"}
  @tens %{0 => "", 1 => "X", 2 => "XX", 3 => "XXX", 4 => "XL", 5 => "L", 6 => "LX", 7 => "LXX", 8 => "LXXX", 9 => "XC"}
  @hundreds %{0 => "", 1 => "C", 2 => "CC", 3 => "CCC", 4 => "CD", 5 => "D", 6 => "DC", 7 => "DCC", 8 => "DCCC", 9 => "CM"}
  @thousands %{0 => "", 1 => "M", 2 => "MM", 3 => "MMM"}
  
  
  @doc """
  Convert the number to a roman number.
  """
  @spec numeral(pos_integer) :: String.t()
  def numeral(number) do
    "#{@thousands[div(number, 1000)]}#{@hundreds[rem(div(number, 100), 10)]}#{@tens[rem(div(number, 10), 10)]}#{@units[rem(number, 10)]}"
  end
end
