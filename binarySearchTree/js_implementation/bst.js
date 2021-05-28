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

    delete(data){
        this.deleteNodeRecursively(this.head, data)
    }

    deleteNodeRecursively(head, data){
        console.log('---del--', head)
        if(head === null){
            return head
        }else if(head.data > data){
            head.right = this.deleteNodeRecursively(head.right, data)
            return head
        }else if(head.data < data) {
            head.left = this.deleteNodeRecursively(head.left, data)
            return head
        }else{
            console.log('----found---', head)
            if(head.left === null && head.right === null){ // leaf node
                head = null
                return head
            }
            if(head.left === null) { // deleting node having only right node
                head = head.right
                return head
            }
            if(head.right === null) { // deleting node having only left node
                head = head.left
                return head
            }
            // deleting node having both children
            const minNode = this.findMinNode(head.right)
            head.data = minNode.data
            head.right = this.deleteNodeRecursively(head.right, minNode.data)
            return head
        }
    }

    minValue(){
        const min = this.findMinNode(this.head)
        console.log(min)
        return min.data
    }
    findMinNode(head){
        // console.log(head.data)
        if(!head) return head
        if(head.left == null){
            return head
        }else{
            return this.findMinNode(head.left)
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
// console.log('---min-----', b.minValue())
b.deleteNodeRecursively(7)
b.inorderTraversal()

