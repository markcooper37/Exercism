pub struct Matcher<T> {
    matcher: Box<dyn Fn(T) -> bool>,
    subs: String,
}

impl<T> Matcher<T> {
    pub fn new<F: Fn(T) -> bool + 'static, S: ToString>(matcher: F, subs: S) -> Matcher<T> {
        Matcher {
            matcher: Box::new(matcher),
            subs: subs.to_string(),
        }
    }
}

pub struct Fizzy<T> {
    matchers: Vec<Matcher<T>>,
}

impl<T: Clone + ToString> Fizzy<T> {
    pub fn new() -> Self {
        Self {
            matchers: Vec::new(),
        }
    }

    #[must_use]
    pub fn add_matcher(mut self, matcher: Matcher<T>) -> Self {
        self.matchers.push(matcher);
        self
    }

    pub fn apply<I: Iterator<Item = T>>(self, iter: I) -> impl Iterator<Item = String> {
        iter.map(move |x| {
            let result: String = self
                .matchers
                .iter()
                .filter(|m| (m.matcher)(x.clone()))
                .map(|m| m.subs.clone())
                .collect();
            if result.is_empty() {
                x.to_string()
            } else {
                result
            }
        })
    }
}

/// convenience function: return a Fizzy which applies the standard fizz-buzz rules
pub fn fizz_buzz<
    T: ToString + PartialEq + From<u8> + Clone + std::ops::Rem<Output = T> + std::fmt::Display,
>() -> Fizzy<T> {
    Fizzy::new()
        .add_matcher(Matcher::new(|x: T| x % 3.into() == 0.into(), "fizz"))
        .add_matcher(Matcher::new(|x: T| x % 5.into() == 0.into(), "buzz"))
}
