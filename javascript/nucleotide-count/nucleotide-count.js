export function countNucleotides(strand) {
  let counts = strand.split('').reduce((counts, letter) => {
    if (!['A', 'C', 'G', 'T'].includes(letter)) {
      throw new Error('Invalid nucleotide in strand');
    };
    counts[letter]++;
    return counts;
  }, {'A': 0, 'C': 0, 'G': 0, 'T': 0});
  return `${counts['A']} ${counts['C']} ${counts['G']} ${counts['T']}`;
}
