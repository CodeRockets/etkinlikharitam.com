select json_agg(sub) venue_json
from (
SELECT
	v.id,
	v.fs_id,
	v.name,
	v.phone,
	v.twitter,
	v.facebook_id,
	v.facebook_username,
	v.url,
	v.city,
	v.country,
	v.formatted_address,
	v.fs_tags,
	v.fs_rating,
	v.fs_rating_signals,
	v.address,
	v.lat,
	v.lon,
	v.bx_tag,
	v.bx_url,
	v.bx_name,
	v.bx_image,
	v.bx_desc,
	v.bx_directions,
	json_agg(e.*) as events
FROM venue v
INNER JOIN event e on (v.id=e.venue_id)
where e.event_date_start BETWEEN $1 and $2
GROUP BY  
	v.id,
	v.fs_id,
	v.name,
	v.phone,
	v.twitter,
	v.facebook_id,
	v.facebook_username,
	v.url,
	v.city,
	v.country,
	v.formatted_address,
	v.fs_tags,
	v.fs_rating,
	v.fs_rating_signals,
	v.address,
	v.lat,
	v.lon,
	v.bx_tag,
	v.bx_url,
	v.bx_name,
	v.bx_image,
	v.bx_desc,
	v.bx_directions
) sub
