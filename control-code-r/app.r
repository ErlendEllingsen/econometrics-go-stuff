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