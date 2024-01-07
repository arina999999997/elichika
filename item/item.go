package item

// this package is for item literals because writing them everytime is painful
// the naming convention is not strict, but ideally prefer english server name
// for unique items like Wish / Radiance, the type is client.Content
// for items with similar one like memorial / memento then it might be a map or a builder or whatever
// after obtaining the client.Content, use Amount() to set the amount, but the amount is defaulted to 1 if not set
