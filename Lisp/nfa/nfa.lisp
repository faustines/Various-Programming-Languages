;;ECS 140A Hw3
;;Faustine Yiu 913062973
;;Source: Lisp Lecture for ECS 140A, Stackoverflow

(defun reachable (transition start final input)
    ;; TODO: Incomplete function

    (cond ((equal (get-length input) 0) (equal start final)))

    (setq next (funcall transition start (car input)))

    (setq map (mapcar transition (list start (car next)) input))
    (setq map (streamline map))

    (setq map2 (mapcar transition (list start (car next) (car(cdr(cdr map)))) input))
    (setq map2 (streamline map2))

    (setq map3 (mapcar transition (list start (car next) (car(cdr(cdr map))) (car(cdr(cdr(cdr map2))))) input))
    (setq map3 (streamline map3))

    (setq map4 (mapcar transition (list start (car next) (car(cdr(cdr map))) (car(cdr(cdr(cdr map2)))) (car(cdr(cdr(cdr(cdr (cdr map3))))))) input))
    (setq map4 (streamline map4))

    (if (null input)
      (equal start final)
    (if (equal (get-length input) 1)
      (cond ((equal (car next) final) t)
      (t (if (equal (car(cdr next)) final) t
        nil)))
    (if (equal (get-length input) 3)
      (if (equal (car(cdr(cdr(cdr map2)))) final)
        t
      (cond ((equal (car(cdr(cdr(cdr(cdr map2))))) final) t)
        (t (setq map2 (mapcar transition (list start (car next) (car(cdr(cdr(cdr map))))) input))
        (setq map2 (streamline map2))
            (cond ((equal (car(cdr(cdr(cdr map2)))) final) t)
              (t (setq map2 (mapcar transition (list start (car next) (car(cdr map))) input))
                 (setq map2 (streamline map2))
                 (if (equal (car(cdr(cdr map2))) final)
                  t
                 nil))))))
    (cond ((equal (get-length input) 2)
      (cond ((equal (car(cdr(cdr map))) final) t)
        (t (setq map (mapcar transition (list start (car (cdr next))) input))
            (setq map (streamline map))
            (cond ((equal (car(cdr(cdr map))) final) t)
              (t nil)))))
    (t
    (cond ((equal (get-length input) 5)
      (cond ((equal (car(cdr(cdr(cdr(cdr(cdr (cdr map4))))))) final) t)
      (t (cond ((equal (car(cdr(cdr(cdr(cdr(cdr(cdr(cdr map4)))))))) final) t)
          (t (setq map3 (mapcar transition (list start (car next) (car(cdr(cdr map))) (car(cdr map2))) input))
             (setq map3 (streamline map3))
             (setq map4 (mapcar transition (list start (car next) (car(cdr(cdr map))) (car(cdr(cdr(cdr map2)))) (car(cdr(cdr(cdr map3))))) input))
             (setq map4 (streamline map4))
          (if (equal (car(cdr(cdr(cdr(cdr map4))))) final)
            t
          nil)))))))
    ))

    )))
  )

(defun get-length (l)
  (if (null l) 0
  (+ 1 (get-length (cdr l)))))

(defun streamline (l)
  (cond ((null l) nil)
        ((atom (car l)) (cons (car l) (streamline (cdr l))))
        (t (append (streamline (car l)) (streamline (cdr l))))))
