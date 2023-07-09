#include <cmath>

// daily_rate calculates the daily rate given an hourly rate
double daily_rate(double hourly_rate) {
    return 8 * hourly_rate;
}

// apply_discount calculates the price after a discount
double apply_discount(double before_discount, double discount) {
    return before_discount * (100 - discount) / 100;
}

// monthly_rate calculates the monthly rate, given an hourly rate and a discount
int monthly_rate(double hourly_rate, double discount) {
    return std::ceil(apply_discount(daily_rate(hourly_rate) * 22, discount));
}

// days_in_budget calculates the number of workdays given a budget, hourly rate,
// and discount
int days_in_budget(int budget, double hourly_rate, double discount) {
    return std::floor(budget / apply_discount(daily_rate(hourly_rate), discount));
}