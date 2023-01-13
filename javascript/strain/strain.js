//
// This is only a SKELETON file for the 'Strain' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const keep = (collection, predicate) => {
  return newCollection(collection, predicate, true);
};

export const discard = (collection, predicate) => {
  return newCollection(collection, predicate, false);
};

export const newCollection = (collection, predicate, decision) => {
  let newCollection = [];
  for (let i = 0; i < collection.length; i++) {
    if (predicate(collection[i]) === decision) {
      newCollection.push(collection[i]);
    };
  };
  return newCollection;
}
