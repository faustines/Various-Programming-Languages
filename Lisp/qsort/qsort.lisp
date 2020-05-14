;;ECS 140A Hw3
;;Faustine Yiu 913062973
;;Source: Lisp Lecture for ECS 140A, Stackoverflow

(defun pivot (n xs)
  (cond ((null xs) (list nil nil))
  (t (setq upper '() lower '())
     (pivot-helper n xs upper lower))))

(defun pivot-helper (n xs upper lower)
  (cond ((>= (car xs) n)
    (setq lower (merge-list lower (list (car xs)))))
  (t (setq upper (merge-list upper (list (car xs))))))

  (cond ((null (cdr xs))
    (list upper lower))
  (t (pivot-helper n (cdr xs) upper lower))))

(defun quicksort (xs)
  (cond ((null xs) nil)
        ((is-sorted xs) xs)
  (t (let* ((piv (car xs)) (x (pivot piv xs)))
      (cond ((and (null (car x)) (null (car (cdr x)))) nil)
          ((null (car x)) (merge-list (list piv) (quicksort (cdr (car (cdr x))))))
          ((null (car (cdr x))) (merge-list (quicksort (car x)) (list piv)))
      (t (setq merge (merge-list (quicksort (car x)) (list piv)))
         (merge-list merge (quicksort (cdr (car (cdr x)))))))))))

(defun is-sorted (xs)
    (cond ((equal (get-length xs) 1) t)
        ((< (car xs) (car (cdr xs))) (is-sorted (cdr xs)))
        (nil)))

(defun get-length (xs)
  (if (null xs) 0
  (+ 1 (get-length (cdr xs)))))

(defun merge-list (m1 m2)
  (if (null m1)
    m2
  (cons (car m1) (merge-list (cdr m1) m2))))
