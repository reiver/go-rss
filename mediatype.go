package rss

// MediaType is the RSS (specific) media-type.
//
// One might want to use it with the HTTP "Accept" request header, or the HTTP "Content-Type" response header.
//
// Note that, prior to the adoption of the "application/rss+xml", non-RSS specific media-types were often used with RSS.
// For example:
//
//	• "application/xml"
//	• "text/xml"
//	• "text/plain"
//
// It is NOT recommended to use these other media-types with RSS.
// And to instead use "application/rss+xml".
//
// The rss.MediaType constant exists to make that easier.
const MediaType = "application/rss+xml"
