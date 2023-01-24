export class CustomSet {
  constructor(elements = []) {
    this.elements = elements;
  }

  empty() {
    return this.elements.length === 0;
  }

  contains(element) {
    return this.elements.includes(element);
  }

  add(element) {
    this.elements.push(element);
    return this;
  }

  subset(set2) {
    return this.elements.every(value => set2.elements.includes(value));
  }

  disjoint(set2) {
    return this.elements.every(value => !set2.elements.includes(value));
  }

  eql(set2) {
    return this.subset(set2) && set2.subset(this);
  }

  union(set2) {
    return new CustomSet(this.elements.concat(set2.elements));
  }

  intersection(set2) {
    return new CustomSet(this.elements.filter(value => set2.contains(value)));
  }

  difference(set2) {
    return new CustomSet(this.elements.filter(value => !set2.contains(value)));
  }
}
