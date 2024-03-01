# Read data from command line arguments
args <- commandArgs(trailingOnly = TRUE)

# Extract data from input 
data <- as.numeric(args)

# Check if data is provided
if (length(data) == 0) {
  stop("Error: No data provided")
}

# Calculate trimmed mean with 5% trimming on both ends
trimmed_mean <- mean(data, trim = 0.05)

# Output trimmed mean
cat(trimmed_mean)
