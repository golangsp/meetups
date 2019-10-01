owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}