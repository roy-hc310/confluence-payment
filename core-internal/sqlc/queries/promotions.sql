-- -- name: CreatePlatform :one
-- insert into platforms (
--     name
-- ) values (
--     $1
-- )
-- returning *;

-- -- name: ListPlatforms :many
-- select * from platforms
-- where deleted_at is null and id < $1
-- order by id desc
-- limit $2;

-- -- name: DetailPlatform :one
-- select * from platforms
-- where deleted_at is null and id = $1;

-- -- name: UpdatePlatform :one
-- update platforms
-- set name = $2
-- where id = $1
-- returning *;


-- insert into promotions (

-- )

-- name: CreatePromotion :one
INSERT INTO promotions (name, promotion_type, code, start_time, end_time, shop_id, usage_quantity, usage_limit_per_user)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;

-- name: CreateProducts :exec
INSERT INTO products (promotion_id, sku, name, purchase_limit)
VALUES (
    unnest($1::integer[]),
    unnest($2::varchar[]),
    unnest($3::varchar[]),
    unnest($4::integer[])
)
RETURNING id;

-- name: CreateProductVariants :exec
INSERT INTO product_variants (product_id, sku, name, discounted_price, discounted_percentage, stock_limit, is_active)
VALUES (
    unnest($1::integer[]),
    unnest($2::varchar[]),
    unnest($3::varchar[]),
    unnest($4::float[]),
    unnest($5::float[]),
    unnest($6::integer[]),
    unnest($7::boolean[])
);