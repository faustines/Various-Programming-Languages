;;ECS 140A Hw3
;;Faustine Yiu 913062973
;;Source: Lisp Lecture for ECS 140A, Stackoverflow, EEL-5840 Notes

(defun matrix-add (m1 m2)
  (mapcar (lambda (a b) (mapcar #'+ a b)) m1 m2))

(defun matrix-multiply (a b)
  (matrix-multiply-helper a (matrix-transpose b)))

(defun matrix-multiply-helper (m1 m2)
  (cond ((null m1) nil)
  (t (cons (mapcar (lambda (a) (dot-product (car m1) a)) m2) (matrix-multiply-helper (cdr m1) m2)))))

(defun dot-product (v1 v2)
  (cond ((null v1) nil) ((null v2) nil)
    ((equal (get-length v1) (get-length v2)) (multiply v1 v2))
    (t nil)))

(defun multiply (v1 v2)
  (cond ((null v1) 0)
  (t (+ (* (car v1) (car v2)) (multiply (cdr v1) (cdr v2))))))

(defun get-length (xs)
  (if (null xs) 0
  (+ 1 (get-length (cdr xs)))))

(defun matrix-transpose (m)
  (cond ((null m) nil)
        ((null (car m)) nil)
        (t (cons (get-first m) (matrix-transpose (get-rest m))))))

(defun get-first (m)
  (if (null m)
      nil
      (cons (car (car m)) (get-first (cdr m)))))

(defun get-rest (m)
  (if (null m)
      nil
      (cons (cdr (car m)) (get-rest (cdr m)))))
