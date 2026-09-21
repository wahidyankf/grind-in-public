from forum_be_python.schemas import PostOut


def build_reply_tree(posts: list[PostOut]) -> list[PostOut]:
    """Attach every post to its parent's replies and return the top-level posts.

    Sibling order follows the input order, so the caller decides it in SQL. The given rows are updated in place.
    """
    by_id = {post.id: post for post in posts}
    roots: list[PostOut] = []
    for post in posts:
        parent = by_id.get(post.parent_id) if post.parent_id is not None else None
        (roots if parent is None else parent.replies).append(post)
    return roots
