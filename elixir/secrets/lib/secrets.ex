defmodule Secrets do
  def secret_add(secret) do
    fn num -> 
      num + secret
    end
  end

  def secret_subtract(secret) do
    fn num ->
      num - secret
    end
  end

  def secret_multiply(secret) do
    fn num ->
      num * secret
    end
  end

  def secret_divide(secret) do
    fn num ->
      trunc(num / secret)
    end
  end

  def secret_and(secret) do
    fn num ->
      Bitwise.band(num, secret)
    end
  end

  def secret_xor(secret) do
    fn num ->
      Bitwise.bxor(num, secret)
    end
  end

  def secret_combine(secret_function1, secret_function2) do
    fn num ->
      secret_function2.(secret_function1.(num))
    end
  end
end
