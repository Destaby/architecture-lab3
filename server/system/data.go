package system

import (
	"database/sql"
)

type Balancer struct {
	Id   	int64  `json:"id"`
	Job 	string `json:"job"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

// func (s *Store) ListBalancers() ([]*Channel, error) {
// 	rows, err := s.Db.Query("SELECT balancers.id, ( SELECT COUNT(*) FROM machines WHERE machines.balancer_id = balancers.id ) as total_machines, machines.id as machine_id FROM balancers INNER JOIN machines ON machines.balancer_id = balancers.id WHERE machines.working = TRUE GROUP BY balancers.id, machines.id LIMIT 200")
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	var res []*Balancer
// 	for rows.Next() {
// 		var c Balancer
// 		if err := rows.Scan(&c.Id, &c.Job); err != nil {
// 			return nil, err
// 		}
// 		res = append(res, &c)
// 	}
// 	if res == nil {
// 		res = make([]*Balancer, 0)
// 	}
// 	return res, nil
// }

type Machine struct {
	Id   				int64  `json:"id"`
	Working 		bool 	 `json:"working"`
	Balancer_id	int64  `json:"id"`
}

func (s *Store) UpdateMachineStatus(id int64, working bool) error {
	_, err := s.Db.Exec("UPDATE machines SET working = $1 where id = $2", working, id)
	return err
}
