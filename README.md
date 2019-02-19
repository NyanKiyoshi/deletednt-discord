# Deletedn't Discord

Deletedn't Discord is a discord experiment where a bot would prefetch 
every message sent to a channel with attachments, and then, 
sent it back using discord's backend cache.

Discord is caching retrieved attachments per IP address for a week or more. 
It thus lets enough time to the bot to retrieve deleted messages attachments 
and send them back to the original channel.
