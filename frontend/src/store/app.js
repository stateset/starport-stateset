module.exports = {
  types: [
    // this line is used by starport scaffolding
		{ type: "contact", fields: ["firstName", "lastName", "email", "phone", "company", ] },
		{ type: "invoice", fields: ["name", "reason", "amountDue", "amountPaid", "total", ] },
		{ type: "purchaseorder", fields: ["number", "name", "type", "status", "startDate", "endDate", "total", "hash", ] },
		{ type: "agreement", fields: ["number", "name", "type", "status", "startDate", "endDate", "total", "hash", ] },
		{ type: "proposal", fields: ["number", "name", "type", "status", "startDate", "endDate", "total", "hash", ] },
		{ type: "user", fields: ["name", "email", ] },
		{ type: "number", fields: ["name", "reason", "amountDue", "amountPaid", "total", ] },
  ],
};
