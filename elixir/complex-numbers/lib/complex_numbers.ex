defmodule ComplexNumbers do
  @typedoc """
  In this module, complex numbers are represented as a tuple-pair containing the real and
  imaginary parts.
  For example, the real number `1` is `{1, 0}`, the imaginary number `i` is `{0, 1}` and
  the complex number `4+3i` is `{4, 3}'.
  """
  @type complex :: {float, float}

  @doc """
  Return the real part of a complex number
  """
  @spec real(a :: complex) :: float
  def real({real, _}), do: real

  @doc """
  Return the imaginary part of a complex number
  """
  @spec imaginary(a :: complex) :: float
  def imaginary({_, imaginary}), do: imaginary

  @doc """
  Multiply two complex numbers, or a real and a complex number
  """
  @spec mul(a :: complex | float, b :: complex | float) :: complex
  def mul({a1, b1}, {a2, b2}), do: {a1 * a2 - b1 * b2, a1 * b2 + a2 * b1}
  def mul({a1, b1}, b), do: mul({a1, b1}, {b, 0})
  def mul(a, {a2, b2}), do: mul({a, 0}, {a2, b2})

  @doc """
  Add two complex numbers, or a real and a complex number
  """
  @spec add(a :: complex | float, b :: complex | float) :: complex
  def add({a1, b1}, {a2, b2}), do: {a1 + a2, b1 + b2}
  def add({a1, b1}, b), do: add({a1, b1}, {b, 0})
  def add(a, {a2, b2}), do: add({a, 0}, {a2, b2})

  @doc """
  Subtract two complex numbers, or a real and a complex number
  """
  @spec sub(a :: complex | float, b :: complex | float) :: complex
  def sub({a1, b1}, {a2, b2}), do: {a1 - a2, b1 - b2}
  def sub({a1, b1}, b), do: sub({a1, b1}, {b, 0})
  def sub(a, {a2, b2}), do: sub({a, 0}, {a2, b2})

  @doc """
  Divide two complex numbers, or a real and a complex number
  """
  @spec div(a :: complex | float, b :: complex | float) :: complex
  def div({a1, b1}, {a2, b2}) do
    {(a1 * a2 + b1 * b2) / (a2 ** 2 + b2 ** 2), (b1 * a2 - a1 * b2) / (a2 ** 2 + b2 ** 2)}
  end
  def div({a1, b1}, b), do: ComplexNumbers.div({a1, b1}, {b, 0})
  def div(a, {a2, b2}), do: ComplexNumbers.div({a, 0}, {a2, b2})

  @doc """
  Absolute value of a complex number
  """
  @spec abs(a :: complex) :: float
  def abs({a, b}), do: :math.sqrt(a ** 2 + b ** 2)

  @doc """
  Conjugate of a complex number
  """
  @spec conjugate(a :: complex) :: complex
  def conjugate({a, b}), do: {a, -b}

  @doc """
  Exponential of a complex number
  """
  @spec exp(a :: complex) :: complex
  def exp({a, b}), do: {:math.exp(a) * :math.cos(b), :math.exp(a) * :math.sin(b)}
end
