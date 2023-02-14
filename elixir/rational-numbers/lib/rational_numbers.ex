defmodule RationalNumbers do
  @type rational :: {integer, integer}

  @doc """
  Add two rational numbers
  """
  @spec add(a :: rational, b :: rational) :: rational
  def add({a1, b1}, {a2, b2}) do
    reduce({a1 * b2 + b1 * a2, b1 * b2})
  end

  @doc """
  Subtract two rational numbers
  """
  @spec subtract(a :: rational, b :: rational) :: rational
  def subtract({a1, b1}, {a2, b2}) do
    reduce({a1 * b2 - b1 * a2, b1 * b2})
  end

  @doc """
  Multiply two rational numbers
  """
  @spec multiply(a :: rational, b :: rational) :: rational
  def multiply({a1, b1}, {a2, b2}) do
    reduce({a1 * a2, b1 * b2})
  end

  @doc """
  Divide two rational numbers
  """
  @spec divide_by(num :: rational, den :: rational) :: rational
  def divide_by({a1, b1}, {a2, b2}) do
    reduce({a1 * b2, b1 * a2})
  end

  @doc """
  Absolute value of a rational number
  """
  @spec abs(a :: rational) :: rational
  def abs({a1, b1}) do
    reduce({Kernel.abs(a1), Kernel.abs(b1)})
  end

  @doc """
  Exponentiation of a rational number by an integer
  """
  @spec pow_rational(a :: rational, n :: integer) :: rational
  def pow_rational({a1, b1}, n) do
    cond do
      n >= 0 -> reduce({Integer.pow(a1, n), Integer.pow(b1, n)})
      true -> reduce({Integer.pow(b1, -n), Integer.pow(a1, -n)})
    end
  end

  @doc """
  Exponentiation of a real number by a rational number
  """
  @spec pow_real(x :: integer, n :: rational) :: float
  def pow_real(x, {n1, n2}) do
    Float.pow(Float.pow(x / 1, n1), 1 / n2)
  end

  @doc """
  Reduce a rational number to its lowest terms
  """
  @spec reduce(a :: rational) :: rational
  def reduce({a1, b1}) do
    gcda = Integer.gcd(a1, b1)
    cond do
      b1 >= 0 -> {trunc(a1 / gcda), trunc(b1 / gcda)}
      true -> {trunc(-a1 / gcda), trunc(-b1 / gcda)}
    end
  end
end
