;;ECS 140A Hw3
;;Faustine Yiu 913062973
;;Source: Lisp Lecture for ECS 140A, Stackoverflow

(defun match (pattern assertion)
   (cond ((equal pattern assertion) t)
        ((and (null pattern) (null assertion)) t)
        ((> (get-length pattern) (get-length assertion)) nil)
    (t (cond ((equal (car pattern) '?)
          (if (> (get-length assertion) 0)
              (if (and (null (cdr pattern)) (null (cdr assertion)))
                  t
              (match (cdr pattern) (cdr assertion)))
          nil))
  ((equal (car pattern) '!)
      (if (> 0 (get-length assertion))
          nil
      (cond ((null (cdr pattern)) t)
      ((or (equal (car (cdr pattern)) '!) (equal (car (cdr pattern)) '?))
      (match (cdr pattern) (cdr assertion)))
      ((not (null (car (cdr pattern))))
        (match-helper pattern assertion)))))
  ((equal (car pattern) (car assertion))
      (if (and (null (cdr pattern)) (null (cdr assertion)))
          (if (equal (get-length pattern) (get-length assertion))
              t
          nil)
      (match (cdr pattern) (cdr assertion))))))))

(defun get-length (xs)
  (if (null xs) 0
  (+ 1 (get-length (cdr xs)))))

(defun match-helper (pattern assertion)
  (cond ((null assertion)
    nil)
  (t (cond ((equal (car (cdr pattern)) (car (cdr assertion)))
    (match (cdr pattern) (cdr assertion)))
    (t (match-helper pattern (cdr assertion)))))))
