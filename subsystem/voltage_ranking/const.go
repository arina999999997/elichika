package voltage_ranking

const VoltageRankingLimit = 50

// duration in seconds
const VoltageRankingCache = 60
const VoltageRankingDeckCache = 60

// technically caching these things separately can lead to the following:
// - cache for voltage ranking is generated
// - someone got a new score
// - someone check out that person's old score
// - but they fetch the new score instead
// it's not too important for now
