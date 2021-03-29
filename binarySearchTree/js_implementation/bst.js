class Node {
    constructor(data) {
        this.data = data
        this.left = null
        this.right = null
    }
}

class BinarySearchTree {
    constructor() {
        this.head = null
    }

    insert(data) {
        const n = new Node(data)
        if (!this.head) {
            this.head = n
        } else {
            this.insertNodeRecursively(this.head, n)
        }
    }

    insertNodeRecursively(head, node) {
        if (head) {
            if (head.data > node.data) {
                if (head.left) {
                    this.insertNodeRecursively(head.left, node)
                } else {
                    head.left = node
                }
            } else if (head.data < node.data) {
                if (head.right) {
                    this.insertNodeRecursively(head.right, node)
                } else {
                    head.right = node
                }
            }
        }
    }

    inorderTraversal() {
        this.recursiveInorderTraversal(this.head)
    }

    recursiveInorderTraversal(head) {
        if (head) {
            this.recursiveInorderTraversal(head.left)
            console.log(head.data)
            this.recursiveInorderTraversal(head.right)
        }
    }
}

const b = new BinarySearchTree()
b.insert(5)
b.insert(3)
b.insert(4)
b.insert(5)
b.insert(9)
b.insert(7)
b.insert(11)
console.log(JSON.stringify(b))
b.inorderTraversal()

