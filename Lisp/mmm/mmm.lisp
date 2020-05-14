;;ECS 140A Hw3
;;Faustine Yiu 913062973
;;Source: Lisp Lecture for ECS 140A, Stackoverflow

(defun min-mean-max (xs)
    (list (get-min xs) (avg xs) (get-max xs)))

(defun avg (l)
    (defun get_sum (l)
        (if (equal l nil) 0
        (+ (car l) (get_sum (cdr l)) )))
    (defun get-length (l)
      (if (null l) 0
      (+ 1 (get-length (cdr l)))))
    (if (equal l nil) 0
        (/ (get_sum l) (get-length l))))

(defun get-min (l)
 (cond ((null (cdr l))(car l))
   ((< (car l) (car(cdr l)))
    (get-min (cons (car l)(cdr (cdr l)))))
   (t
    (get-min (cdr l)))))

(defun get-max (l)
 (cond ((null (cdr l))(car l))
   ((> (car l) (car(cdr l)))
    (get-max (cons (car l)(cdr (cdr l)))))
   (t
    (get-max (cdr l)))))
