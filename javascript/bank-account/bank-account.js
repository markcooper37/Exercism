export class BankAccount {
  constructor() {
    [this._open, this._balance] = [false, 0];
  }

  open() {
    if (this._open) {
      throw new ValueError();
    }
    this._open= true;
  }

  close() {
    if (!this._open) {
      throw new ValueError();
    }
    [this._open, this._balance] = [false, 0];
  }

  deposit(amount) {
    if (!this._open || amount < 0) {
      throw new ValueError();
    }
    this._balance += amount;
  }

  withdraw(amount) {
    if (!this._open || amount < 0 || amount > this._balance) {
      throw new ValueError();
    }
    this._balance -= amount;
  }

  get balance() {
    if (!this._open) {
      throw new ValueError();
    }
    return this._balance;
  }
}

export class ValueError extends Error {
  constructor() {
    super('Bank account error');
  }
}
