/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
 var reverseList = function(rest) {
    if(!rest){ // list doesnot exists
        return rest
    }
    
    if(!rest.next){ // only one element
        return rest
    }
    
    let first = reverseList(rest.next)  // divide list into two parts: rest & first 
    rest.next.next = rest   //  make previous next next value point to previous
    rest.next = null 
    return first
};
