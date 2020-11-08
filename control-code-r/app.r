x <- c(5, 8, 9, 14, 17)
y <- c(10, 12, 15, 20, 21)

# Statistical properties
stats <- list(
  x = list(
    mean = mean(x),
    sd = sd(x)
  ),
  y = list(
    mean = mean(y),
    sd = sd(y)
  )
)

# Regression
reg <- lm(y ~ x)
regSum <- summary(reg)

rss <- sum(regSum$residuals^2)

# R^2
R2 = regSum$r.squared
# Adjusted R^2
R2_adj = regSum$adj.r.squared
