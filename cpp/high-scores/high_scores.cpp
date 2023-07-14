#include "high_scores.h"

#include <algorithm>

namespace arcade {

    std::vector<int> HighScores::list_scores() {
        return scores;
    }

    int HighScores::latest_score() {
        return scores.back();
    }

    int HighScores::personal_best() {
        return *max_element(scores.begin(), scores.end());
    }

    std::vector<int> HighScores::top_three() {
        std::vector<int> best = scores;
        sort(best.begin(), best.end(), std::greater<>());
        return {best.begin(), best.begin() + std::min(3, int(best.size()))};
    }

}  // namespace arcade
