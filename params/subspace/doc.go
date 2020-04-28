package subspace

/*
To prevent namespace collision between consumer modules, we define type
"space". A Space can only be generated by the keeper, and the keeper checks
the existence of the space having the same name before generating the
space.

Consumer modules must take a space (via Keeper.Subspace), not the keeper
itself. This isolates each modules from the others and make them modify the
parameters safely. Keeper can be treated as master permission for all
subspaces (via Keeper.GetSubspace), so should be passed to proper modules
(ex. gov)
*/