#[derive(Debug, PartialEq)]
pub struct CustomSet<T> {
    elements: Vec<T>,
}

impl<T: Copy + PartialEq + Ord> CustomSet<T> {
    pub fn new(_input: &[T]) -> Self {
        let mut elements: Vec<T> = _input.to_vec();
        elements.sort();
        Self { elements: elements }
    }

    pub fn contains(&self, _element: &T) -> bool {
        self.elements.contains(_element)
    }

    pub fn add(&mut self, _element: T) {
        if !self.contains(&_element) {
            self.elements.push(_element);
            self.elements.sort();
        }
    }

    pub fn is_subset(&self, _other: &Self) -> bool {
        self.elements.iter().all(|x| _other.contains(x))
    }

    pub fn is_empty(&self) -> bool {
        self.elements.is_empty()
    }

    pub fn is_disjoint(&self, _other: &Self) -> bool {
        self.intersection(_other).is_empty()
    }

    #[must_use]
    pub fn intersection(&self, _other: &Self) -> Self {
        Self {
            elements: self
                .elements
                .clone()
                .into_iter()
                .filter(|x| _other.contains(x))
                .collect(),
        }
    }

    #[must_use]
    pub fn difference(&self, _other: &Self) -> Self {
        Self {
            elements: self
                .elements
                .clone()
                .into_iter()
                .filter(|x| !(_other.contains(x)))
                .collect(),
        }
    }

    #[must_use]
    pub fn union(&self, _other: &Self) -> Self {
        let mut new_set = Self {
            elements: self.elements.clone(),
        };
        for element in _other.elements.clone() {
            new_set.add(element)
        }
        new_set
    }
}
